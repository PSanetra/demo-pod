import { Component } from '@angular/core';
import { HttpHeadersService } from './http-headers.service';
import { Observable } from 'rxjs';
import { map, tap } from 'rxjs/operators';

@Component({
  selector: 'app-http-headers',
  templateUrl: './http-headers.component.html',
  styleUrls: ['./http-headers.component.scss']
})
export class HttpHeadersComponent {

  httpHeadersList$: Observable<{name: string; value: string}[]> = this.httpHeadersService.getHttpHeadersInfo().pipe(
    map(m => Object.keys(m).map((k: string) => ({ name: k, value: m[k] })))
  );

  constructor(private httpHeadersService: HttpHeadersService) {
  }

}
