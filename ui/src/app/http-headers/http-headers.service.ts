import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { HttpHeadersInfo } from './http-headers-info';
import { apiUri } from '../../_shared/functions/api-uri';

@Injectable({
  providedIn: 'root'
})
export class HttpHeadersService {
  constructor(private http: HttpClient) {
  }

  getHttpHeadersInfo(): Observable<HttpHeadersInfo> {
    return this.http.get<HttpHeadersInfo>(`${apiUri()}/http/headers`);
  }
}
