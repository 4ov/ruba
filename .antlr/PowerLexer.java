// Generated from /home/mo/opnlab/v2/Power.g4 by ANTLR 4.8
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class PowerLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.8", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, Ident=6, String=7, Number=8, NORMALSTRING=9, 
		WS=10, Int=11, HEX=12, Float=13, HEX_FLOAT=14;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"T__0", "T__1", "T__2", "T__3", "T__4", "Ident", "String", "Number", 
			"NORMALSTRING", "WS", "Int", "HEX", "Float", "HEX_FLOAT", "ExponentPart", 
			"HexExponentPart", "EscapeSequence", "DecimalEscape", "HexEscape", "UtfEscape", 
			"Digit", "HexDigit"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'const'", "'<'", "'>'", "'let'", "'='"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, "Ident", "String", "Number", "NORMALSTRING", 
			"WS", "Int", "HEX", "Float", "HEX_FLOAT"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}


	public PowerLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "Power.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\20\u00f0\b\1\4\2"+
		"\t\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4"+
		"\13\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22"+
		"\t\22\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\3\2\3\2\3\2\3"+
		"\2\3\2\3\2\3\3\3\3\3\4\3\4\3\5\3\5\3\5\3\5\3\6\3\6\3\7\3\7\7\7B\n\7\f"+
		"\7\16\7E\13\7\3\b\3\b\3\t\3\t\5\tK\n\t\3\n\3\n\3\n\7\nP\n\n\f\n\16\nS"+
		"\13\n\3\n\3\n\3\13\6\13X\n\13\r\13\16\13Y\3\13\3\13\3\f\6\f_\n\f\r\f\16"+
		"\f`\3\r\3\r\3\r\6\rf\n\r\r\r\16\rg\3\16\6\16k\n\16\r\16\16\16l\3\16\3"+
		"\16\7\16q\n\16\f\16\16\16t\13\16\3\16\5\16w\n\16\3\16\3\16\6\16{\n\16"+
		"\r\16\16\16|\3\16\5\16\u0080\n\16\3\16\6\16\u0083\n\16\r\16\16\16\u0084"+
		"\3\16\3\16\5\16\u0089\n\16\3\17\3\17\3\17\6\17\u008e\n\17\r\17\16\17\u008f"+
		"\3\17\3\17\7\17\u0094\n\17\f\17\16\17\u0097\13\17\3\17\5\17\u009a\n\17"+
		"\3\17\3\17\3\17\3\17\6\17\u00a0\n\17\r\17\16\17\u00a1\3\17\5\17\u00a5"+
		"\n\17\3\17\3\17\3\17\6\17\u00aa\n\17\r\17\16\17\u00ab\3\17\3\17\5\17\u00b0"+
		"\n\17\3\20\3\20\5\20\u00b4\n\20\3\20\6\20\u00b7\n\20\r\20\16\20\u00b8"+
		"\3\21\3\21\5\21\u00bd\n\21\3\21\6\21\u00c0\n\21\r\21\16\21\u00c1\3\22"+
		"\3\22\3\22\3\22\5\22\u00c8\n\22\3\22\3\22\3\22\3\22\5\22\u00ce\n\22\3"+
		"\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\5\23\u00db\n\23"+
		"\3\24\3\24\3\24\3\24\3\24\3\25\3\25\3\25\3\25\3\25\6\25\u00e7\n\25\r\25"+
		"\16\25\u00e8\3\25\3\25\3\26\3\26\3\27\3\27\2\2\30\3\3\5\4\7\5\t\6\13\7"+
		"\r\b\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35\20\37\2!\2#\2%\2\'\2)\2"+
		"+\2-\2\3\2\16\5\2C\\aac|\6\2\62;C\\aac|\4\2$$^^\5\2\13\f\16\17\"\"\4\2"+
		"ZZzz\4\2GGgg\4\2--//\4\2RRrr\f\2$$))^^cdhhppttvvxx||\3\2\62\64\3\2\62"+
		";\5\2\62;CHch\2\u010a\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2"+
		"\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25"+
		"\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\3/\3\2\2"+
		"\2\5\65\3\2\2\2\7\67\3\2\2\2\t9\3\2\2\2\13=\3\2\2\2\r?\3\2\2\2\17F\3\2"+
		"\2\2\21J\3\2\2\2\23L\3\2\2\2\25W\3\2\2\2\27^\3\2\2\2\31b\3\2\2\2\33\u0088"+
		"\3\2\2\2\35\u00af\3\2\2\2\37\u00b1\3\2\2\2!\u00ba\3\2\2\2#\u00cd\3\2\2"+
		"\2%\u00da\3\2\2\2\'\u00dc\3\2\2\2)\u00e1\3\2\2\2+\u00ec\3\2\2\2-\u00ee"+
		"\3\2\2\2/\60\7e\2\2\60\61\7q\2\2\61\62\7p\2\2\62\63\7u\2\2\63\64\7v\2"+
		"\2\64\4\3\2\2\2\65\66\7>\2\2\66\6\3\2\2\2\678\7@\2\28\b\3\2\2\29:\7n\2"+
		"\2:;\7g\2\2;<\7v\2\2<\n\3\2\2\2=>\7?\2\2>\f\3\2\2\2?C\t\2\2\2@B\t\3\2"+
		"\2A@\3\2\2\2BE\3\2\2\2CA\3\2\2\2CD\3\2\2\2D\16\3\2\2\2EC\3\2\2\2FG\5\23"+
		"\n\2G\20\3\2\2\2HK\5\27\f\2IK\5\33\16\2JH\3\2\2\2JI\3\2\2\2K\22\3\2\2"+
		"\2LQ\7$\2\2MP\5#\22\2NP\n\4\2\2OM\3\2\2\2ON\3\2\2\2PS\3\2\2\2QO\3\2\2"+
		"\2QR\3\2\2\2RT\3\2\2\2SQ\3\2\2\2TU\7$\2\2U\24\3\2\2\2VX\t\5\2\2WV\3\2"+
		"\2\2XY\3\2\2\2YW\3\2\2\2YZ\3\2\2\2Z[\3\2\2\2[\\\b\13\2\2\\\26\3\2\2\2"+
		"]_\5+\26\2^]\3\2\2\2_`\3\2\2\2`^\3\2\2\2`a\3\2\2\2a\30\3\2\2\2bc\7\62"+
		"\2\2ce\t\6\2\2df\5-\27\2ed\3\2\2\2fg\3\2\2\2ge\3\2\2\2gh\3\2\2\2h\32\3"+
		"\2\2\2ik\5+\26\2ji\3\2\2\2kl\3\2\2\2lj\3\2\2\2lm\3\2\2\2mn\3\2\2\2nr\7"+
		"\60\2\2oq\5+\26\2po\3\2\2\2qt\3\2\2\2rp\3\2\2\2rs\3\2\2\2sv\3\2\2\2tr"+
		"\3\2\2\2uw\5\37\20\2vu\3\2\2\2vw\3\2\2\2w\u0089\3\2\2\2xz\7\60\2\2y{\5"+
		"+\26\2zy\3\2\2\2{|\3\2\2\2|z\3\2\2\2|}\3\2\2\2}\177\3\2\2\2~\u0080\5\37"+
		"\20\2\177~\3\2\2\2\177\u0080\3\2\2\2\u0080\u0089\3\2\2\2\u0081\u0083\5"+
		"+\26\2\u0082\u0081\3\2\2\2\u0083\u0084\3\2\2\2\u0084\u0082\3\2\2\2\u0084"+
		"\u0085\3\2\2\2\u0085\u0086\3\2\2\2\u0086\u0087\5\37\20\2\u0087\u0089\3"+
		"\2\2\2\u0088j\3\2\2\2\u0088x\3\2\2\2\u0088\u0082\3\2\2\2\u0089\34\3\2"+
		"\2\2\u008a\u008b\7\62\2\2\u008b\u008d\t\6\2\2\u008c\u008e\5-\27\2\u008d"+
		"\u008c\3\2\2\2\u008e\u008f\3\2\2\2\u008f\u008d\3\2\2\2\u008f\u0090\3\2"+
		"\2\2\u0090\u0091\3\2\2\2\u0091\u0095\7\60\2\2\u0092\u0094\5-\27\2\u0093"+
		"\u0092\3\2\2\2\u0094\u0097\3\2\2\2\u0095\u0093\3\2\2\2\u0095\u0096\3\2"+
		"\2\2\u0096\u0099\3\2\2\2\u0097\u0095\3\2\2\2\u0098\u009a\5!\21\2\u0099"+
		"\u0098\3\2\2\2\u0099\u009a\3\2\2\2\u009a\u00b0\3\2\2\2\u009b\u009c\7\62"+
		"\2\2\u009c\u009d\t\6\2\2\u009d\u009f\7\60\2\2\u009e\u00a0\5-\27\2\u009f"+
		"\u009e\3\2\2\2\u00a0\u00a1\3\2\2\2\u00a1\u009f\3\2\2\2\u00a1\u00a2\3\2"+
		"\2\2\u00a2\u00a4\3\2\2\2\u00a3\u00a5\5!\21\2\u00a4\u00a3\3\2\2\2\u00a4"+
		"\u00a5\3\2\2\2\u00a5\u00b0\3\2\2\2\u00a6\u00a7\7\62\2\2\u00a7\u00a9\t"+
		"\6\2\2\u00a8\u00aa\5-\27\2\u00a9\u00a8\3\2\2\2\u00aa\u00ab\3\2\2\2\u00ab"+
		"\u00a9\3\2\2\2\u00ab\u00ac\3\2\2\2\u00ac\u00ad\3\2\2\2\u00ad\u00ae\5!"+
		"\21\2\u00ae\u00b0\3\2\2\2\u00af\u008a\3\2\2\2\u00af\u009b\3\2\2\2\u00af"+
		"\u00a6\3\2\2\2\u00b0\36\3\2\2\2\u00b1\u00b3\t\7\2\2\u00b2\u00b4\t\b\2"+
		"\2\u00b3\u00b2\3\2\2\2\u00b3\u00b4\3\2\2\2\u00b4\u00b6\3\2\2\2\u00b5\u00b7"+
		"\5+\26\2\u00b6\u00b5\3\2\2\2\u00b7\u00b8\3\2\2\2\u00b8\u00b6\3\2\2\2\u00b8"+
		"\u00b9\3\2\2\2\u00b9 \3\2\2\2\u00ba\u00bc\t\t\2\2\u00bb\u00bd\t\b\2\2"+
		"\u00bc\u00bb\3\2\2\2\u00bc\u00bd\3\2\2\2\u00bd\u00bf\3\2\2\2\u00be\u00c0"+
		"\5+\26\2\u00bf\u00be\3\2\2\2\u00c0\u00c1\3\2\2\2\u00c1\u00bf\3\2\2\2\u00c1"+
		"\u00c2\3\2\2\2\u00c2\"\3\2\2\2\u00c3\u00c4\7^\2\2\u00c4\u00ce\t\n\2\2"+
		"\u00c5\u00c7\7^\2\2\u00c6\u00c8\7\17\2\2\u00c7\u00c6\3\2\2\2\u00c7\u00c8"+
		"\3\2\2\2\u00c8\u00c9\3\2\2\2\u00c9\u00ce\7\f\2\2\u00ca\u00ce\5%\23\2\u00cb"+
		"\u00ce\5\'\24\2\u00cc\u00ce\5)\25\2\u00cd\u00c3\3\2\2\2\u00cd\u00c5\3"+
		"\2\2\2\u00cd\u00ca\3\2\2\2\u00cd\u00cb\3\2\2\2\u00cd\u00cc\3\2\2\2\u00ce"+
		"$\3\2\2\2\u00cf\u00d0\7^\2\2\u00d0\u00db\5+\26\2\u00d1\u00d2\7^\2\2\u00d2"+
		"\u00d3\5+\26\2\u00d3\u00d4\5+\26\2\u00d4\u00db\3\2\2\2\u00d5\u00d6\7^"+
		"\2\2\u00d6\u00d7\t\13\2\2\u00d7\u00d8\5+\26\2\u00d8\u00d9\5+\26\2\u00d9"+
		"\u00db\3\2\2\2\u00da\u00cf\3\2\2\2\u00da\u00d1\3\2\2\2\u00da\u00d5\3\2"+
		"\2\2\u00db&\3\2\2\2\u00dc\u00dd\7^\2\2\u00dd\u00de\7z\2\2\u00de\u00df"+
		"\5-\27\2\u00df\u00e0\5-\27\2\u00e0(\3\2\2\2\u00e1\u00e2\7^\2\2\u00e2\u00e3"+
		"\7w\2\2\u00e3\u00e4\7}\2\2\u00e4\u00e6\3\2\2\2\u00e5\u00e7\5-\27\2\u00e6"+
		"\u00e5\3\2\2\2\u00e7\u00e8\3\2\2\2\u00e8\u00e6\3\2\2\2\u00e8\u00e9\3\2"+
		"\2\2\u00e9\u00ea\3\2\2\2\u00ea\u00eb\7\177\2\2\u00eb*\3\2\2\2\u00ec\u00ed"+
		"\t\f\2\2\u00ed,\3\2\2\2\u00ee\u00ef\t\r\2\2\u00ef.\3\2\2\2 \2CJOQY`gl"+
		"rv|\177\u0084\u0088\u008f\u0095\u0099\u00a1\u00a4\u00ab\u00af\u00b3\u00b8"+
		"\u00bc\u00c1\u00c7\u00cd\u00da\u00e8\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}