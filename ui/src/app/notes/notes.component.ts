import { Component, OnDestroy } from '@angular/core';
import { FormControl } from '@angular/forms';
import { debounceTime, map, shareReplay, startWith, switchMap, takeUntil, tap } from 'rxjs/operators';
import { NotesService } from './notes.service';
import { componentDestroyed, OnDestroyMixin } from '@w11k/ngx-componentdestroyed';
import { Observable } from 'rxjs';

@Component({
  selector: 'app-notes',
  templateUrl: './notes.component.html',
  styleUrls: ['./notes.component.scss']
})
export class NotesComponent extends OnDestroyMixin implements OnDestroy {

  control$: Observable<FormControl> = this.notesService.getNotes().pipe(
    map(v => new FormControl(v)),
    // start with dummy form control to prevent error: Cannot find control with unspecified name attribute
    startWith(new FormControl('')),
    shareReplay({bufferSize: 1, refCount: true})
  );

  constructor(private notesService: NotesService) {
    super();

    this.control$.pipe(
      switchMap(c => c.valueChanges),
      debounceTime(250),
      takeUntil(componentDestroyed(this)),
      switchMap(notes => notesService.putNotes(notes))
    ).subscribe();

  }

  ngOnDestroy(): void {
  }

}
