// 枚举都按4字节对齐
{$Z4+}

type
  // 重定定义
  TUnixDateTime = type LongInt; // int64改为LongInt。只精确到秒
  TPngImage = type TPortableNetworkGraphic;
  TMonthCalendar = type TCalendar;
  TJPEGPixelFormat = TPixelFormat;
  TGauge = class(TGraphicControl);
  // TStringHelper = Type Helper for AnsiString
  // TGuidHelper = record helper for TGUID

  //TListItemsHelper = class Helper for TListItems

  // treeview.inc 补丁
  // procedure TTreeNodes.DefineProperties(Filer: TFiler);
  // Filer.DefineBinaryProperty('NodeData', @ReadData, @WriteData, WriteNodes);

  // listitems.inc 补丁
  // procedure TListItems.DefineProperties(Filer: TFiler);
  // Filer.DefineBinaryProperty('ItemData', @ReadData, nil, false);

