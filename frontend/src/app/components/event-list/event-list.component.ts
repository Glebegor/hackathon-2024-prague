import { Component, Input } from '@angular/core';
import { ISubscribtion } from '../../services/events.service';
import { RouterLink } from '@angular/router';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-event-list',
  standalone: true,
  imports: [CommonModule, RouterLink],
  templateUrl: './event-list.component.html',
  styleUrl: './event-list.component.scss'
})
export class EventListComponent {
  @Input() title:string = "Channel list";
  @Input() subtitle:string;
  @Input() subscribtions: ISubscribtion[];
}
