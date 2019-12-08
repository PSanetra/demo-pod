import { Routes } from '@angular/router';

export const routes: Routes = [
  {
    path: '',
    pathMatch: 'full',
    loadChildren: () => import('./landing-page/landing-page.module').then(m => m.LandingPageModule)
  },
  {
    path: 'cpu',
    pathMatch: 'full',
    loadChildren: () => import('./cpu/cpu.module').then(m => m.CpuModule)
  },
  {
    path: 'env',
    pathMatch: 'full',
    loadChildren: () => import('./env/env.module').then(m => m.EnvModule)
  },
  {
    path: 'ip',
    pathMatch: 'full',
    loadChildren: () => import('./ip/ip.module').then(m => m.IpModule)
  },
  {
    path: 'liveness',
    pathMatch: 'full',
    loadChildren: () => import('./liveness/liveness.module').then(m => m.LivenessModule)
  },
  {
    path: 'notes',
    pathMatch: 'full',
    loadChildren: () => import('./notes/notes.module').then(m => m.NotesModule)
  },
  {
    path: 'memory',
    pathMatch: 'full',
    loadChildren: () => import('./memory/memory.module').then(m => m.MemoryModule)
  },
  {
    path: 'readiness',
    pathMatch: 'full',
    loadChildren: () => import('./readiness/readiness.module').then(m => m.ReadinessModule)
  },
  {
    path: 'watch-files',
    pathMatch: 'full',
    loadChildren: () => import('./watch-files/watch-files.module').then(m => m.WatchFilesModule)
  },
  {
    path: '**',
    loadChildren: () => import('./not-found/not-found.module').then(m => m.NotFoundModule)
  }
];
