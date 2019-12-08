import { NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';
import { MatFormFieldModule } from '@angular/material/form-field';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { LivenessComponent } from './liveness.component';
import { routes } from './liveness.routes';
import { MatCardModule } from '@angular/material/card';
import { MatSlideToggleModule } from '@angular/material/slide-toggle';
import { LivenessDescriptionPipe } from './liveness-description.pipe';
import { LatestControlValuePipeModule } from '../../_shared/pipes/latest-control-value/latest-control-value-pipe.module';

@NgModule({
  declarations: [
    LivenessComponent,
    LivenessDescriptionPipe
  ],
  imports: [
    CommonModule,
    RouterModule.forChild(routes),
    MatFormFieldModule,
    MatCardModule,
    MatSlideToggleModule,
    FormsModule,
    ReactiveFormsModule,
    LatestControlValuePipeModule
  ],
  providers: [],
})
export class LivenessModule {
}
