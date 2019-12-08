import { NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';
import { MatFormFieldModule } from '@angular/material/form-field';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { MatCardModule } from '@angular/material/card';
import { LatestControlValuePipeModule } from '../../_shared/pipes/latest-control-value/latest-control-value-pipe.module';
import { NgxMaterialTimepickerModule } from 'ngx-material-timepicker';
import { MatInputModule } from '@angular/material/input';
import { ReadinessComponent } from './readiness.component';
import { ReadinessDescriptionPipe } from './readiness-description.pipe';
import { routes } from './readiness.routes';

@NgModule({
  declarations: [
    ReadinessComponent,
    ReadinessDescriptionPipe
  ],
  imports: [
    CommonModule,
    RouterModule.forChild(routes),
    MatFormFieldModule,
    MatCardModule,
    MatFormFieldModule,
    MatInputModule,
    NgxMaterialTimepickerModule,
    FormsModule,
    ReactiveFormsModule,
    LatestControlValuePipeModule,
  ],
  providers: [],
})
export class ReadinessModule {
}
