import { Component, Input } from '@angular/core';
import { ISubscribtion } from '../../services/events.service';
import { RouterLink } from '@angular/router';
import { CommonModule } from '@angular/common';
import { debounceTime, distinctUntilChanged, map } from 'rxjs/operators';
import { BehaviorSubject } from 'rxjs/internal/BehaviorSubject';
import { Observable } from 'rxjs/internal/Observable';

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
  private filterSubject: BehaviorSubject<string> = new BehaviorSubject<string>('');
  filter$: Observable<string> = this.filterSubject.asObservable();

  filteredSubscriptions$: Observable<any[]> = this.filter$.pipe(
    debounceTime(300),
    distinctUntilChanged(),
    map((filterValue: string) => {
      if (!filterValue.trim()) {
        return this.subscribtions;
      }
      return this.subscribtions.filter(sub => {
        // Фильтруем по названию и хэштегам
        const nameMatch = sub.name.toLowerCase().includes(filterValue.toLowerCase());
        const tagMatch = sub.tags.toLowerCase().includes(filterValue.toLowerCase());
        return nameMatch || tagMatch;
      });
    })
  );

  onFilterChanged(filterValue: string): void {
    this.filterSubject.next(filterValue);
  }
}
