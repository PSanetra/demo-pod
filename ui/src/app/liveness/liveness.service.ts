import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from '../../environments/environment';
import { map } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class LivenessService {
  constructor(private http: HttpClient) {
  }

  getLiveness(): Observable<boolean> {
    return this.http.get<{alive: boolean}>(`${environment.apiUri}/liveness`).pipe(
      map(r => r.alive)
    );
  }

  putLiveness(alive: boolean): Observable<void> {
    return this.http.put<void>(`${environment.apiUri}/liveness`, {
      alive: alive
    });
  }
}
