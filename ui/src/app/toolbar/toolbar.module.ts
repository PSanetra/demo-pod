import { NgModule } from '@angular/core';
import { ToolbarComponent } from './toolbar.component';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatButtonModule } from '@angular/material/button';
import { RouterModule } from '@angular/router';

@NgModule({
  declarations: [
    ToolbarComponent
  ],
  imports: [
    MatToolbarModule,
    MatButtonModule,
    RouterModule
  ],
  exports: [
    ToolbarComponent
  ],
  providers: [],
})
export class ToolbarModule {
}
