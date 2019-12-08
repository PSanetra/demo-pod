import { Routes } from '@angular/router';
import { CpuComponent } from './cpu.component';

export const routes: Routes = [
  {
    path: '',
    pathMatch: 'full',
    component: CpuComponent
  }
];
