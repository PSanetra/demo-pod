import { Component } from '@angular/core';
import { Observable } from 'rxjs';
import { map, shareReplay, tap } from 'rxjs/operators';
import { WatchFilesService } from './watch-files.service';

interface WatchedFile {
  name: string;
  content$: Observable<string>;
}

@Component({
  selector: 'app-watch-files',
  templateUrl: './watch-files.component.html',
  styleUrls: ['./watch-files.component.scss']
})
export class WatchFilesComponent {

  watchedFiles$: Observable<WatchedFile[]> = this.watchFilesService.getWhitelistedFiles().pipe(
    map(files => files.map(file => ({
        name: file,
        content$: this.watchFilesService.watchFile(file).pipe(
          shareReplay({bufferSize: 1, refCount: true})
        )
      })
    )),
    shareReplay({bufferSize: 1, refCount: true})
  );

  constructor(private watchFilesService: WatchFilesService) {
  }

}
