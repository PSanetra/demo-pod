import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'readinessDescription',
  pure: true
})
export class ReadinessDescriptionPipe implements PipeTransform {
    transform(time: unknown): string {
      if (typeof time !== 'string') {
        return '';
      }

      const now = new Date();
      const hourMatch = /(\d+):\d+/.exec(time);
      const minuteMatch = /\d+:(\d+)/.exec(time);
      const hour = hourMatch && (+hourMatch[1]) || 0;
      const minute = minuteMatch && (+minuteMatch[1]) || 0;

      if (hour > now.getUTCHours()) {
        return '/api/ready returns 500 - Internal Server Error';
      } else if (hour === now.getUTCHours() && minute > now.getUTCMinutes()) {
        return '/api/ready returns 500 - Internal Server Error';
      } else {
        return '/api/ready returns 204 - No Content';
      }
    }
}
