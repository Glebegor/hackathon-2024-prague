import { Component, OnInit, inject,ViewChild } from '@angular/core';
import { TelegramService } from '../../services/telegram.service';
import { CommonModule } from '@angular/common';
import { EventsService } from '../../services/events.service';
import { EventListComponent } from '../../components/event-list/event-list.component';
import { HeaderComponent } from '../../components/header/header.component';



@Component({
  selector: 'app-events',
  standalone: true,
  imports: [CommonModule,EventListComponent,HeaderComponent],
  templateUrl: './events.component.html',
  styleUrl: './events.component.scss',
  providers: [HeaderComponent] // Ensure HeaderComponent is provided
})
export class EventsComponent{
  //inject btn
  telegram= inject(TelegramService);
  subscribtions = inject(EventsService);

  @ViewChild(EventListComponent) eventListComponent: EventListComponent;


  constructor(){
    this.telegram.MainButton.show();
    this.telegram.BackButton.hide();
    this.telegram.MainButton.setText("Add channel");

  }
  
  onFilterChanged(filterValue: string): void {
    this.eventListComponent.onFilterChanged(filterValue);
  }
}
