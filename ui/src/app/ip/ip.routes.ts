import { Routes } from '@angular/router';
import { IpComponent } from './ip.component';

export const routes: Routes = [
  {
    path: '',
    pathMatch: 'full',
    component: IpComponent
  }
];
