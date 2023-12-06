import { NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';
import { HttpHeadersComponent } from './http-headers.component';
import { routes } from './http-headers.routes';
import { CommonModule } from '@angular/common';
import { MatTableModule } from '@angular/material/table';

@NgModule({
  declarations: [
    HttpHeadersComponent
  ],
  imports: [
    CommonModule,
    RouterModule.forChild(routes),
    MatTableModule,
  ],
  exports: [
    HttpHeadersComponent
  ],
  providers: [],
})
export class HttpHeadersModule {
}
