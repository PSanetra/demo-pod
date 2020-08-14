import { Component, OnDestroy } from '@angular/core';
import { Observable } from 'rxjs';
import { map, shareReplay, startWith, switchMap, takeUntil } from 'rxjs/operators';
import { CpuService } from './cpu.service';
import { FormControl } from '@angular/forms';
import { componentDestroyed, OnDestroyMixin } from '@w11k/ngx-componentdestroyed';

@Component({
  selector: 'app-cpu',
  templateUrl: './cpu.component.html',
  styleUrls: ['./cpu.component.scss']
})
export class CpuComponent extends OnDestroyMixin implements OnDestroy{

  stress$: Observable<number> = this.cpuService.getCpuStress().pipe(shareReplay({bufferSize: 1, refCount: true}));
  utilization$: Observable<number[]> = this.cpuService.getCpuUtilization().pipe(shareReplay({bufferSize: 1, refCount: true}));

  threadsControl$: Observable<FormControl> = this.stress$.pipe(
    map(s => new FormControl(s)),
    startWith(new FormControl()),
    shareReplay({bufferSize: 1, refCount: true})
  );

  constructor(private cpuService: CpuService) {
    super();

    this.threadsControl$.pipe(
      switchMap(c => c.valueChanges),
      switchMap(threads => this.cpuService.putCpuStress(threads)),
      takeUntil(componentDestroyed(this))
    ).subscribe();

  }

  ngOnDestroy(): void {
  }

}
