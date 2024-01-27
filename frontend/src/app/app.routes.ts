import { Routes } from '@angular/router';
import { EventComponent } from './pages/event/event.component';
import { EventsComponent } from './pages/events/events.component';

export const routes: Routes = [
    {path: '',component: EventsComponent, pathMatch:'full'},
    {path: 'event/:id',component: EventComponent},
];
