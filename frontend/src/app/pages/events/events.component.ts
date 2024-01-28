import { Component, OnInit, inject } from '@angular/core';
import { TelegramService } from '../../services/telegram.service';
import { CommonModule } from '@angular/common';
import { EventsService } from '../../services/events.service';
import { EventListComponent } from '../../components/event-list/event-list.component';
import { HeaderComponent } from '../../components/header/header.component';
import { CacheService } from '../../services/cache.service';

@Component({
  selector: 'app-events',
  standalone: true,
  imports: [CommonModule,EventListComponent,HeaderComponent],
  templateUrl: './events.component.html',
  styleUrl: './events.component.scss'
})
export class EventsComponent{
  //inject btn
  telegram= inject(TelegramService);
  subscribtions = inject(EventsService);

  constructor(){
    this.telegram.MainButton.show();
    this.telegram.BackButton.hide();
    this.telegram.MainButton.setText("Add channel");
  
  }

}
