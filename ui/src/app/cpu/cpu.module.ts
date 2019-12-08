import { NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';
import { CommonModule } from '@angular/common';
import { CpuComponent } from './cpu.component';
import { routes } from './cpu.routes';
import { MatCardModule } from '@angular/material/card';
import { MatTableModule } from '@angular/material/table';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MatInputModule } from '@angular/material/input';

@NgModule({
  declarations: [
    CpuComponent
  ],
  imports: [
    CommonModule,
    RouterModule.forChild(routes),
    MatCardModule,
    MatTableModule,
    MatInputModule,
    FormsModule,
    ReactiveFormsModule
  ],
  exports: [
    CpuComponent
  ],
  providers: [],
})
export class CpuModule {
}
