import { NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';
import { IpComponent } from './ip.component';
import { routes } from './ip.routes';
import { CommonModule } from '@angular/common';
import { MatTableModule } from '@angular/material/table';

@NgModule({
  declarations: [
    IpComponent
  ],
  imports: [
    CommonModule,
    RouterModule.forChild(routes),
    MatTableModule
  ],
  exports: [
    IpComponent
  ],
  providers: [],
})
export class IpModule {
}
