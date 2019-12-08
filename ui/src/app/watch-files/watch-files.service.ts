import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from '../../environments/environment';
import { ServerSentEventsService } from '../../_shared/services/sse/server-sent-events.service';
import { map } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class WatchFilesService {
  constructor(private http: HttpClient, private sse: ServerSentEventsService) {
  }

  getWhitelistedFiles(): Observable<string[]> {
    return this.http.get<string[]>(`${environment.apiUri}/watch-whitelist`);
  }

  watchFile(file: string): Observable<string> {
    return this.sse.get<{content: string}>(`${environment.apiUri}/watch/${file}?mode=sse`).pipe(
      map(e => e.content)
    );
  }

}
