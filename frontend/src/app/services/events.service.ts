import { Injectable } from '@angular/core';
import { group } from '@angular/animations';

export enum IEventsGroup {
  events = "events",
}
export interface IEvent {
  id: number;
  name: string;
  idOfPerson: string;
  price: number;
  description: string;
  image: string;
  tags:string;
  startDate: string;
  category: IEventsGroup;
  link:string;
}
const events: IEvent[] = [
  {
    id: 1,
    name: "falafel",
    idOfPerson: "124142sdfsfd",
    price:69,
    description:
      "Your perfect pack for everyday use and walks in the forest. Stash your laptop (up to 15 inches) in the padded sleeve, your everyday",
    category: IEventsGroup.events,
    image: "https://fakestoreapi.com/img/71-3HjGNDUL._AC_SY879._SX._UX._SY._UY_.jpg",
    tags:"#bom",
    startDate:"1.1.2022",
    link:"http://google.com"
  },
  {
    id: 2, 
    name: "falafel",
    idOfPerson: "124142sdfsfd",
    price:69,
    description:
      "Your perfect pack for everyday use and walks in the forest. Stash your laptop (up to 15 inches) in the padded sleeve, your everyday",
    category: IEventsGroup.events,
    image: "https://fakestoreapi.com/img/81fPKd-2AYL._AC_SL1500_.jpg",
    tags:"#bom",
    startDate:"1.1.2022",
    link:"http://google.com"
  },
  {
    id: 3,
    name: "falafel",
    idOfPerson: "124142sdfsfd",
    price:69,
    description:
      "Your perfect pack for everyday use and walks in the forest. Stash your laptop (up to 15 inches) in the padded sleeve, your everyday",
    category: IEventsGroup.events,
    image: "https://fakestoreapi.com/img/71-3HjGNDUL._AC_SY879._SX._UX._SY._UY_.jpg",
    tags:"#bom",
    startDate:"1.1.2022",
    link:"http://google.com"
  },
  {
    id: 4, 
    name: "falafel",
    idOfPerson: "124142sdfsfd",
    price:69,
    description:
      "Your perfect pack for everyday use and walks in the forest. Stash your laptop (up to 15 inches) in the padded sleeve, your everyday",
    category: IEventsGroup.events,
    image: "https://fakestoreapi.com/img/81fPKd-2AYL._AC_SL1500_.jpg",
    tags:"#bom",
    startDate:"1.1.2022",
    link:"http://google.com"
  },
  {
    id: 5,
    name: "falafel",
    idOfPerson: "124142sdfsfd",
    price:69,
    description:
      "Your perfect pack for everyday use and walks in the forest. Stash your laptop (up to 15 inches) in the padded sleeve, your everyday",
    category: IEventsGroup.events,
    image: "https://fakestoreapi.com/img/71-3HjGNDUL._AC_SY879._SX._UX._SY._UY_.jpg",
    tags:"#bom",
    startDate:"1.1.2022",
    link:"http://google.com"
  },
  {
    id: 6, 
    name: "falafel",
    idOfPerson: "124142sdfsfd",
    price:69,
    description:
      "Your perfect pack for everyday use and walks in the forest. Stash your laptop (up to 15 inches) in the padded sleeve, your everyday",
    category: IEventsGroup.events,
    image: "https://fakestoreapi.com/img/81fPKd-2AYL._AC_SL1500_.jpg",
    tags:"#bom",
    startDate:"1.1.2022",
    link:"http://google.com"
  },
  {
    id: 7,
    name: "falafel",
    idOfPerson: "124142sdfsfd",
    price:69,
    description:
      "Your perfect pack for everyday use and walks in the forest. Stash your laptop (up to 15 inches) in the padded sleeve, your everyday",
    category: IEventsGroup.events,
    image: "https://fakestoreapi.com/img/71-3HjGNDUL._AC_SY879._SX._UX._SY._UY_.jpg",
    tags:"#bom",
    startDate:"1.1.2022",
    link:"http://google.com"
  },
  {
    id: 8, 
    name: "falafel",
    idOfPerson: "124142sdfsfd",
    price:69,
    description:
      "Your perfect pack for everyday use and walks in the forest. Stash your laptop (up to 15 inches) in the padded sleeve, your everyday",
    category: IEventsGroup.events,
    image: "https://fakestoreapi.com/img/81fPKd-2AYL._AC_SL1500_.jpg",
    tags:"#bom",
    startDate:"1.1.2022",
    link:"http://google.com"
  },
];

``
@Injectable({
  providedIn: 'root'
})
export class EventsService {
  readonly events: IEvent[] = events;

  getById(id: number) {  
    return this.events.find(product => product.id === id);
  }

  get byGroup() {
    return this.events.reduce((group: { [key: string]: IEvent[] }, event) => {
      if (!group[event.category]) {
        group[event.category] = [];
      }
      group[event.category].push(event);
      return group;
    }, {});
  }


}
