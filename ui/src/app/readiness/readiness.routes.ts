import { Routes } from '@angular/router';
import { ReadinessComponent } from './readiness.component';

export const routes: Routes = [
  {
    path: '',
    pathMatch: 'full',
    component: ReadinessComponent
  }
];
