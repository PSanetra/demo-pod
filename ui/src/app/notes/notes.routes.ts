import { Routes } from '@angular/router';
import { NotesComponent } from './notes.component';

export const routes: Routes = [
  {
    path: '',
    pathMatch: 'full',
    component: NotesComponent
  }
];
