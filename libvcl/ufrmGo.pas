unit ufrmGo;

interface

uses
  Winapi.Windows, Winapi.Messages, System.SysUtils, System.Variants, System.Classes, Vcl.Graphics,
  Vcl.Controls, Vcl.Forms;

type
  TGoForm = class(TForm)
  protected
    procedure InitializeNewForm; override;
  end;

implementation

{$R *.dfm}

var
  // 默认设置为False，但会显示为True，原因估计是加载
  // dfm文件问题
  uGlobalFormScaled: Boolean = False;


procedure SetGlobalFormScaled(AValue: LongBool); stdcall;
begin
  uGlobalFormScaled := AValue;
end;

procedure Form_ScaleForPPI(AObj: TGoForm; ANewPPI: Integer); stdcall;
begin
  AObj.ScaleForPPI(ANewPPI);
end;

procedure Form_ScaleControlsForDpi(AObj: TGoForm; ANewPPI: Integer); stdcall;
begin
  AObj.ScaleControlsForDpi(ANewPPI);
end;


{ TGoForm }

procedure TGoForm.InitializeNewForm;
begin
  inherited InitializeNewForm;
  Scaled := uGlobalFormScaled;
  if not Scaled then
    PixelsPerInch := 96;
end;

exports
   SetGlobalFormScaled,
   Form_ScaleForPPI,
   Form_ScaleControlsForDpi;



end.
