import { Component, OnDestroy } from '@angular/core';
import { Observable } from 'rxjs';
import { map, shareReplay, startWith, switchMap, takeUntil } from 'rxjs/operators';
import { FormControl } from '@angular/forms';
import { componentDestroyed } from '@w11k/ngx-componentdestroyed';
import { MemoryService } from './memory.service';
import { MemoryUtilization } from './memory-utilization';

@Component({
  selector: 'app-memory',
  templateUrl: './memory.component.html',
  styleUrls: ['./memory.component.scss']
})
export class MemoryComponent implements OnDestroy{

  block$: Observable<number> = this.memoryService.getMemoryBlock().pipe(shareReplay({bufferSize: 1, refCount: true}));
  utilization$: Observable<MemoryUtilization> = this.memoryService.getMemoryUtilization().pipe(
    shareReplay({bufferSize: 1, refCount: true})
  );

  blockControl$: Observable<FormControl> = this.block$.pipe(
    map(s => new FormControl(s)),
    startWith(new FormControl()),
    shareReplay({bufferSize: 1, refCount: true})
  );

  constructor(private memoryService: MemoryService) {

    this.blockControl$.pipe(
      switchMap(c => c.valueChanges),
      switchMap(bytes => this.memoryService.putMemoryBlock(bytes)),
      takeUntil(componentDestroyed(this))
    ).subscribe();

  }

  ngOnDestroy(): void {
  }

}
