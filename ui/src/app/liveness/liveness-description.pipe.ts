import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'livenessDescription',
  pure: true
})
export class LivenessDescriptionPipe implements PipeTransform {
    transform(alive: unknown): string {
      if (alive) {
        return '/api/alive returns 204 - No Content';
      } else {
        return '/api/alive returns 500 - Internal Server Error';
      }
    }
}
