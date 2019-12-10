import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { EnvInfo } from './env-info';
import { apiUri } from '../../_shared/functions/api-uri';

@Injectable({
  providedIn: 'root'
})
export class EnvService {
  constructor(private http: HttpClient) {
  }

  getEnvInfo(): Observable<EnvInfo> {
    return this.http.get<EnvInfo>(`${apiUri()}/environment`);
  }
}
