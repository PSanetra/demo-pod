import { Component, OnDestroy } from '@angular/core';
import { FormControl } from '@angular/forms';
import { debounceTime, map, shareReplay, startWith, switchMap, takeUntil, tap } from 'rxjs/operators';
import { componentDestroyed } from '@w11k/ngx-componentdestroyed';
import { Observable } from 'rxjs';
import { LivenessService } from './liveness.service';

@Component({
  selector: 'app-liveness',
  templateUrl: './liveness.component.html',
  styleUrls: ['./liveness.component.scss']
})
export class LivenessComponent implements OnDestroy {

  control$: Observable<FormControl> = this.livenessService.getLiveness().pipe(
    map(v => new FormControl(v)),
    // start with dummy form control to prevent error: Cannot find control with unspecified name attribute
    startWith(new FormControl(true)),
    shareReplay({bufferSize: 1, refCount: true})
  );

  constructor(private livenessService: LivenessService) {

    this.control$.pipe(
      switchMap(c => c.valueChanges),
      debounceTime(250),
      takeUntil(componentDestroyed(this)),
      switchMap(alive => livenessService.putLiveness(alive))
    ).subscribe();

  }

  ngOnDestroy(): void {
  }

}
