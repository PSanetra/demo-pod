import { Component, OnDestroy } from '@angular/core';
import { FormControl } from '@angular/forms';
import { debounceTime, map, shareReplay, startWith, switchMap, takeUntil, tap } from 'rxjs/operators';
import { componentDestroyed, OnDestroyMixin } from '@w11k/ngx-componentdestroyed';
import { Observable } from 'rxjs';
import { ReadinessService } from './readiness.service';

@Component({
  selector: 'app-readiness',
  templateUrl: './readiness.component.html',
  styleUrls: ['./readiness.component.scss']
})
export class ReadinessComponent extends OnDestroyMixin implements OnDestroy {

  control$: Observable<FormControl> = this.readinessService.getReadiness().pipe(
    map(v => new FormControl(v)),
    // start with dummy form control to prevent error: Cannot find control with unspecified name attribute
    startWith(new FormControl('')),
    shareReplay({bufferSize: 1, refCount: true})
  );

  constructor(private readinessService: ReadinessService) {
    super();

    this.control$.pipe(
      switchMap(c => c.valueChanges),
      debounceTime(250),
      takeUntil(componentDestroyed(this)),
      switchMap(readyAfter => readinessService.putReadiness(readyAfter))
    ).subscribe();

  }

  ngOnDestroy(): void {
  }

}
