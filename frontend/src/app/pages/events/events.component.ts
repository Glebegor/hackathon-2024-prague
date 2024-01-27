import { Component, inject } from '@angular/core';
import { TelegramService } from '../../services/telegram.service';
import { CommonModule } from '@angular/common';
import { EventsService } from '../../services/events.service';
import { EventListComponent } from '../../components/event-list/event-list.component';

@Component({
  selector: 'app-events',
  standalone: true,
  imports: [CommonModule,EventListComponent],
  templateUrl: './events.component.html',
  styleUrl: './events.component.scss'
})
export class EventsComponent {
  //inject btn
  telegram= inject(TelegramService);
  events = inject(EventsService);

  constructor(){
    this.telegram.MainButton.show();
    this.telegram.BackButton.hide();
    this.telegram.MainButton.setText("Add Event");

  }
}
