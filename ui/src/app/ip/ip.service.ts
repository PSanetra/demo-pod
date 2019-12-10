import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { IpInfo } from './ip-info';
import { apiUri } from '../../_shared/functions/api-uri';

@Injectable({
  providedIn: 'root'
})
export class IpService {
  constructor(private http: HttpClient) {
  }

  getIpInfo(): Observable<IpInfo> {
    return this.http.get<IpInfo>(`${apiUri()}/ip`);
  }
}
