import { Component, OnDestroy, OnInit } from '@angular/core';
import { EventsService, ISubscribtion,  } from '../../services/events.service';
import { TelegramService } from '../../services/telegram.service';
import { ActivatedRoute, Router } from '@angular/router';
import { CommonModule } from '@angular/common';
import { CacheService } from '../../services/cache.service';

@Component({
  selector: 'app-event',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './event.component.html',
  styleUrl: './event.component.scss'
})
export class EventComponent implements OnInit, OnDestroy {
  subscribtion:ISubscribtion;
  

  constructor(private subscribtions:EventsService, 
    private telegram:TelegramService,
    private route:ActivatedRoute,
    private router:Router,
    private cacheService: CacheService
    ){
      const id = this.route.snapshot.paramMap.get('id');
      this.subscribtion = this.subscribtions.getById(Number(id));
      this.telegram.MainButton.show();
      this.telegram.MainButton.setText("Subscribe!");
      this.goBack = this.goBack.bind(this)
      this.telegram.BackButton.show();
   }
  
  ngOnInit(): void {
    this.telegram.BackButton.onClick(this.goBack);           
  }
  ngOnDestroy(): void {
    this.telegram.BackButton.offClick(this.goBack);
    this.cacheService.reloadPage();
  }
  goBack(){
    this.router.navigate(['/']);
  }
  

}
