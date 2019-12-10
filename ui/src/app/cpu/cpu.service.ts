import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { interval, Observable } from 'rxjs';
import { switchMap } from 'rxjs/operators';
import { apiUri } from '../../_shared/functions/api-uri';

@Injectable({
  providedIn: 'root'
})
export class CpuService {
  constructor(private http: HttpClient) {
  }

  getCpuStress(): Observable<number> {
    return this.http.get<number>(`${apiUri()}/cpu-stress`);
  }

  putCpuStress(threads: number): Observable<void> {
    return this.http.put<void>(`${apiUri()}/cpu-stress`, threads);
  }

  getCpuUtilization(): Observable<number[]> {
    return interval(2000).pipe(
      switchMap(() => this.http.get<number[]>(`${apiUri()}/cpu-utilization`))
    );
  }
}
