import { Routes } from '@angular/router';
import { MemoryComponent } from './memory.component';

export const routes: Routes = [
  {
    path: '',
    pathMatch: 'full',
    component: MemoryComponent
  }
];
