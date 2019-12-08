import { Component } from '@angular/core';
import { IpService } from './ip.service';
import { Observable } from 'rxjs';
import { map, shareReplay } from 'rxjs/operators';
import { IpInfo } from './ip-info';

@Component({
  selector: 'app-ip',
  templateUrl: './ip.component.html',
  styleUrls: ['./ip.component.scss']
})
export class IpComponent {

  ipInfo$: Observable<IpInfo> = this.ipService.getIpInfo().pipe(shareReplay({bufferSize: 1, refCount: true}));

  clientIpList$: Observable<{ type: string, ip: string }[]> = this.ipInfo$.pipe(
    map(info => ([
      {type: 'Client IP', ip: info.client_ip},
      {type: 'Original Client IP', ip: info.original_ip},
    ]))
  );

  podIpList$: Observable<string[]> = this.ipInfo$.pipe(map(info => info.pod_ip_list));

  constructor(private ipService: IpService) {
  }

}
