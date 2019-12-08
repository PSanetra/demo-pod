import { NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';
import { EnvComponent } from './env.component';
import { routes } from './env.routes';
import { CommonModule } from '@angular/common';
import { MatTableModule } from '@angular/material/table';

@NgModule({
  declarations: [
    EnvComponent
  ],
  imports: [
    CommonModule,
    RouterModule.forChild(routes),
    MatTableModule,
  ],
  exports: [
    EnvComponent
  ],
  providers: [],
})
export class EnvModule {
}
