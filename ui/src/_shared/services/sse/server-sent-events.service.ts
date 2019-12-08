import { Injectable, NgZone } from '@angular/core';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ServerSentEventsService {

  constructor(private ngZone: NgZone) {
  }

  get<T>(url: string): Observable<T> {

    return new Observable<T>(observer => {

      const eventSrc = new EventSource(url, {
        withCredentials: true
      });

      const processEvent = (e: MessageEvent) => {
        this.ngZone.run(() => {
          observer.next(JSON.parse(e.data));
        });
      };

      eventSrc.addEventListener('initial', processEvent);
      eventSrc.addEventListener('file-modified', processEvent);

      eventSrc.addEventListener('error', (err: any) => {

        if (err.target.readyState !== EventSource.CLOSED) {
          return;
        }

        observer.error(err);
      });

      return () => {
        eventSrc.close();
      };
    });

  }

}
