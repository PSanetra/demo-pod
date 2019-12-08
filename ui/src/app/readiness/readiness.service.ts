import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from '../../environments/environment';
import { map } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class ReadinessService {
  constructor(private http: HttpClient) {
  }

  getReadiness(): Observable<string> {
    return this.http.get<{ready_after: string}>(`${environment.apiUri}/readiness`).pipe(
      map(r => {
        const match = /\d\d:\d\d/.exec(r.ready_after);

        if (!match || match.length < 1) {
          throw new Error('Could not parse time from ' + r.ready_after);
        }

        return match[0];
      })
    );
  }

  putReadiness(time: string): Observable<void> {
    let readyAfter = JSON.stringify(new Date()).replace(/\d\d:\d\d:\d\d/, time + ':00')
    readyAfter = readyAfter.replace(/"/g, '');

    return this.http.put<void>(`${environment.apiUri}/readiness`, {
      ready_after: readyAfter
    });
  }
}
