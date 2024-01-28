import { Component, OnDestroy, OnInit } from '@angular/core';
import { EventsService, ISubscribtion,  } from '../../services/events.service';
import { TelegramService } from '../../services/telegram.service';
import { ActivatedRoute, Router } from '@angular/router';
import { CommonModule } from '@angular/common';

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
    
    ){
      const id = this.route.snapshot.paramMap.get('id');
      this.subscribtion = this.subscribtions.getById(Number(id));
      this.telegram.MainButton.show();
      this.telegram.MainButton.setText("Subscribe!");
      this.goBack = this.goBack.bind(this)
  }
  goBack(){
    this.router.navigate(['/'])
  }
  ngOnDestroy(): void {
    this.telegram.BackButton.offClick(this.goBack);
  }
  ngOnInit(): void {
    this.telegram.BackButton.show();
    this.telegram.BackButton.onClick(this.goBack);

  }

}
