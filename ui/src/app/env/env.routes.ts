import { Routes } from '@angular/router';
import { EnvComponent } from './env.component';

export const routes: Routes = [
  {
    path: '',
    pathMatch: 'full',
    component: EnvComponent
  }
];
