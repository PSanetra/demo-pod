import { Routes } from '@angular/router';
import { LivenessComponent } from './liveness.component';

export const routes: Routes = [
  {
    path: '',
    pathMatch: 'full',
    component: LivenessComponent
  }
];
