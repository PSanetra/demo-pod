import { Pipe, PipeTransform } from '@angular/core';
import { EMPTY, Observable } from 'rxjs';
import { FormControl } from '@angular/forms';
import { startWith } from 'rxjs/operators';

@Pipe({
  name: 'latestValue$',
  pure: true
})
export class LatestControlValuePipe implements PipeTransform {
  transform<T>(control: FormControl | null): Observable<unknown> {
    if (control == null) {
      return EMPTY;
    }

    return control.valueChanges.pipe(
      startWith(control.value as unknown)
    );
  }
}
