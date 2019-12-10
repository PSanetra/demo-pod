import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { ServerSentEventsService } from '../../_shared/services/sse/server-sent-events.service';
import { map } from 'rxjs/operators';
import { apiUri } from '../../_shared/functions/api-uri';

@Injectable({
  providedIn: 'root'
})
export class WatchFilesService {
  constructor(private http: HttpClient, private sse: ServerSentEventsService) {
  }

  getWhitelistedFiles(): Observable<string[]> {
    return this.http.get<string[]>(`${apiUri()}/watch-whitelist`);
  }

  watchFile(file: string): Observable<string> {
    return this.sse.get<{content: string}>(`${apiUri()}/watch/${file}?mode=sse`).pipe(
      map(e => e.content)
    );
  }

}
