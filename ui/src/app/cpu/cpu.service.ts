import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { interval, Observable } from 'rxjs';
import { environment } from '../../environments/environment';
import { shareReplay, switchMap } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class CpuService {
  constructor(private http: HttpClient) {
  }

  getCpuStress(): Observable<number> {
    return this.http.get<number>(`${environment.apiUri}/cpu-stress`);
  }

  putCpuStress(threads: number): Observable<void> {
    return this.http.put<void>(`${environment.apiUri}/cpu-stress`, threads);
  }

  getCpuUtilization(): Observable<number[]> {
    return interval(2000).pipe(
      switchMap(() => this.http.get<number[]>(`${environment.apiUri}/cpu-utilization`))
    );
  }
}
