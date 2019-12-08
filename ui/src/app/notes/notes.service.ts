import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { Observable, of, throwError } from 'rxjs';
import { environment } from '../../environments/environment';
import { catchError, map } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class NotesService {
  constructor(private http: HttpClient) {
  }

  getNotes(): Observable<string> {
    return this.http.get<{content: string}>(`${environment.apiUri}/notes`).pipe(
      map(r => r.content),
      catchError(e => {
        if (e instanceof HttpErrorResponse && e.status === 404) {
          return of('');
        } else {
          return throwError(e);
        }
      })
    );
  }

  putNotes(notes: string): Observable<void> {
    return this.http.put<void>(`${environment.apiUri}/notes`, {
      content: notes
    });
  }
}
