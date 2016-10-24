import { Pipe, PipeTransform, Injectable } from '@angular/core'

@Pipe({
	name: 'formatPhone'
})
@Injectable()
export class PhonePipe implements PipeTransform {
	transform(value: any, args?: any): any {
		if (value) {
			return '+48 '+value.substring(0,3)+' '+value.substring(3,6)+' '+value.substring(6,9);
		}
	}
}
