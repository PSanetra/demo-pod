import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { interval, Observable } from 'rxjs';
import { environment } from '../../environments/environment';
import { map, switchMap } from 'rxjs/operators';
import { MemoryUtilization } from './memory-utilization';

@Injectable({
  providedIn: 'root'
})
export class MemoryService {
  constructor(private http: HttpClient) {
  }

  getMemoryBlock(): Observable<number> {
    return this.http.get<number>(`${environment.apiUri}/memory-block`);
  }

  putMemoryBlock(bytes: number): Observable<void> {
    return this.http.put<void>(`${environment.apiUri}/memory-block`, bytes);
  }

  getMemoryUtilization(): Observable<MemoryUtilization> {
    return interval(2000).pipe(
      switchMap(() => this.http.get<{in_use: string, available: string}>(`${environment.apiUri}/memory`)),
      map(v => ({inUse: v.in_use, available: v.available}))
    );
  }
}
