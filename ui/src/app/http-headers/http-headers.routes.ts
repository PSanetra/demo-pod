import { Routes } from '@angular/router';
import { HttpHeadersComponent } from './http-headers.component';

export const routes: Routes = [
  {
    path: '',
    pathMatch: 'full',
    component: HttpHeadersComponent
  }
];
