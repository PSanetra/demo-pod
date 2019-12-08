import { Component } from '@angular/core';
import { EnvService } from './env.service';
import { Observable } from 'rxjs';
import { map, tap } from 'rxjs/operators';

@Component({
  selector: 'app-env',
  templateUrl: './env.component.html',
  styleUrls: ['./env.component.scss']
})
export class EnvComponent {

  envList$: Observable<{name: string; value: string}[]> = this.envService.getEnvInfo().pipe(
    map(m => Object.keys(m).map((k: string) => ({ name: k, value: m[k] })))
  );

  constructor(private envService: EnvService) {
  }

}
