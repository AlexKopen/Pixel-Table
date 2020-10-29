import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SettingsPickerComponent } from './settings-picker.component';

describe('SettingsPickerComponent', () => {
  let component: SettingsPickerComponent;
  let fixture: ComponentFixture<SettingsPickerComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [SettingsPickerComponent]
    }).compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(SettingsPickerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
