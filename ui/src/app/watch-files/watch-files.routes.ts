import { Routes } from '@angular/router';
import { WatchFilesComponent } from './watch-files.component';

export const routes: Routes = [
  {
    path: '',
    pathMatch: 'full',
    component: WatchFilesComponent
  }
];
