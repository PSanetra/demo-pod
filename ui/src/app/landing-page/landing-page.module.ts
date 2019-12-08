import { NgModule } from '@angular/core';
import { LandingPageComponent } from './landing-page.component';
import { MatCardModule } from '@angular/material/card';
import { RouterModule } from '@angular/router';
import { routes } from './landine-page.routes';
import { MatIconModule } from '@angular/material/icon';

@NgModule({
  declarations: [
    LandingPageComponent
  ],
  imports: [
    MatCardModule,
    RouterModule.forChild(routes),
    MatCardModule,
    MatIconModule
  ],
  exports: [
    LandingPageComponent
  ],
  providers: [],
})
export class LandingPageModule {
}
