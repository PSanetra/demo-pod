import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { apiUri } from '../../_shared/functions/api-uri';

@Injectable({
  providedIn: 'root'
})
export class LivenessService {
  constructor(private http: HttpClient) {
  }

  getLiveness(): Observable<boolean> {
    return this.http.get<{alive: boolean}>(`${apiUri()}/liveness`).pipe(
      map(r => r.alive)
    );
  }

  putLiveness(alive: boolean): Observable<void> {
    return this.http.put<void>(`${apiUri()}/liveness`, {
      alive: alive
    });
  }
}
