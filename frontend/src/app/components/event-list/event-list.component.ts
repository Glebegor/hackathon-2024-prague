import { Component, Input } from '@angular/core';
import { IEvent } from '../../services/events.service';
import { RouterLink } from '@angular/router';

@Component({
  selector: 'app-event-list',
  standalone: true,
  imports: [RouterLink],
  templateUrl: './event-list.component.html',
  styleUrl: './event-list.component.scss'
})
export class EventListComponent {
  @Input() title:string;
  @Input() subtitle:string;
  @Input() events: IEvent[];
}
