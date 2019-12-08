import { NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';
import { CommonModule } from '@angular/common';
import { routes } from './watch-files.routes';
import { WatchFilesComponent } from './watch-files.component';
import { MatCardModule } from '@angular/material/card';

@NgModule({
  declarations: [
    WatchFilesComponent
  ],
  imports: [
    CommonModule,
    RouterModule.forChild(routes),
    MatCardModule
  ],
  exports: [
    WatchFilesComponent
  ],
  providers: [],
})
export class WatchFilesModule {
}
