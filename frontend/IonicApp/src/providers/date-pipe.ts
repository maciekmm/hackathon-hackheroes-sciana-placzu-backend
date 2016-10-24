import { Pipe, PipeTransform, Injectable } from '@angular/core'

@Pipe({
	name: 'formatDate'
})
@Injectable()
export class DatePipe implements PipeTransform {
	transform(value: any, args?: any): any {
		if (value) {
			var dateParts = value.split("-");
			return dateParts[2]+'.'+dateParts[1]+'.'+dateParts[0];
		}
	}
}
