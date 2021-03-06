// Generated by tabtoy
// Version: 2.8.6
// DO NOT EDIT!!
using System.Collections.Generic;

namespace cfg
{
	
	

	// Defined in table: DataConfig
	
	public partial class DataConfig
	{
	
		public tabtoy.Logger TableLogger = new tabtoy.Logger();
	
		
		/// <summary> 
		/// Equip
		/// </summary>
		public List<EquipDefine> Equip = new List<EquipDefine>(); 
		
		/// <summary> 
		/// Item
		/// </summary>
		public List<ItemDefine> Item = new List<ItemDefine>(); 
	
	
		#region Index code
	 	Dictionary<int, EquipDefine> _EquipByID = new Dictionary<int, EquipDefine>();
        public EquipDefine GetEquipByID(int ID, EquipDefine def = default(EquipDefine))
        {
            EquipDefine ret;
            if ( _EquipByID.TryGetValue( ID, out ret ) )
            {
                return ret;
            }
			
			if ( def == default(EquipDefine) )
			{
				TableLogger.ErrorLine("GetEquipByID failed, ID: {0}", ID);
			}

            return def;
        }
		Dictionary<int, ItemDefine> _ItemByID = new Dictionary<int, ItemDefine>();
        public ItemDefine GetItemByID(int ID, ItemDefine def = default(ItemDefine))
        {
            ItemDefine ret;
            if ( _ItemByID.TryGetValue( ID, out ret ) )
            {
                return ret;
            }
			
			if ( def == default(ItemDefine) )
			{
				TableLogger.ErrorLine("GetItemByID failed, ID: {0}", ID);
			}

            return def;
        }
		
	
		#endregion
		#region Deserialize code
		
		static tabtoy.DeserializeHandler<DataConfig> DataConfigDeserializeHandler = new tabtoy.DeserializeHandler<DataConfig>(Deserialize);
		public static void Deserialize( DataConfig ins, tabtoy.DataReader reader )
		{
 			int tag = -1;
            while ( -1 != (tag = reader.ReadTag()))
            {
                switch (tag)
                { 
                	case 0xa0000:
                	{
						ins.Equip.Add( reader.ReadStruct<EquipDefine>(EquipDefineDeserializeHandler) );
                	}
                	break; 
                	case 0xa0001:
                	{
						ins.Item.Add( reader.ReadStruct<ItemDefine>(ItemDefineDeserializeHandler) );
                	}
                	break; 
                }
             }

			
			// Build Equip Index
			for( int i = 0;i< ins.Equip.Count;i++)
			{
				var element = ins.Equip[i];
				
				ins._EquipByID.Add(element.ID, element);
				
			}
			
			// Build Item Index
			for( int i = 0;i< ins.Item.Count;i++)
			{
				var element = ins.Item[i];
				
				ins._ItemByID.Add(element.ID, element);
				
			}
			
		}
		static tabtoy.DeserializeHandler<EquipDefine> EquipDefineDeserializeHandler = new tabtoy.DeserializeHandler<EquipDefine>(Deserialize);
		public static void Deserialize( EquipDefine ins, tabtoy.DataReader reader )
		{
 			int tag = -1;
            while ( -1 != (tag = reader.ReadTag()))
            {
                switch (tag)
                { 
                	case 0x10000:
                	{
						ins.ID = reader.ReadInt32();
                	}
                	break; 
                	case 0x60001:
                	{
						ins.Name = reader.ReadString();
                	}
                	break; 
                	case 0x60002:
                	{
						ins.Desc = reader.ReadString();
                	}
                	break; 
                }
             }

			
		}
		static tabtoy.DeserializeHandler<ItemDefine> ItemDefineDeserializeHandler = new tabtoy.DeserializeHandler<ItemDefine>(Deserialize);
		public static void Deserialize( ItemDefine ins, tabtoy.DataReader reader )
		{
 			int tag = -1;
            while ( -1 != (tag = reader.ReadTag()))
            {
                switch (tag)
                { 
                	case 0x10000:
                	{
						ins.ID = reader.ReadInt32();
                	}
                	break; 
                	case 0x60001:
                	{
						ins.Name = reader.ReadString();
                	}
                	break; 
                	case 0x60002:
                	{
						ins.Desc = reader.ReadString();
                	}
                	break; 
                }
             }

			
		}
		#endregion
	

	} 

	// Defined in table: Equip
	[System.Serializable]
	public partial class EquipDefine
	{
	
		
		/// <summary> 
		/// 唯一ID
		/// </summary>
		public int ID = 0; 
		
		/// <summary> 
		/// 名字
		/// </summary>
		public string Name = ""; 
		
		/// <summary> 
		/// 描述
		/// </summary>
		public string Desc = ""; 
	
	

	} 

	// Defined in table: Item
	[System.Serializable]
	public partial class ItemDefine
	{
	
		
		/// <summary> 
		/// 唯一ID
		/// </summary>
		public int ID = 0; 
		
		/// <summary> 
		/// 名字
		/// </summary>
		public string Name = ""; 
		
		/// <summary> 
		/// 描述
		/// </summary>
		public string Desc = ""; 
	
	

	} 

}
