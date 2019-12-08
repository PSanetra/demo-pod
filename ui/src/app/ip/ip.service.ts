import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from '../../environments/environment';
import { IpInfo } from './ip-info';

@Injectable({
  providedIn: 'root'
})
export class IpService {
  constructor(private http: HttpClient) {
  }

  getIpInfo(): Observable<IpInfo> {
    return this.http.get<IpInfo>(`${environment.apiUri}/ip`);
  }
}
